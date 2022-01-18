package crawler

import (
	"context"
	"errors"
	"github.com/Elderly-AI/observer/crawler/internal/pkg/model"
	"github.com/Elderly-AI/observer/crawler/internal/pkg/vk"
	"io"
	"net/http"
	"sync"
)

const MaxVkUsersCountRequest = 75
const MaxVkUserRequestBatchSize = 25

// TODO обернуть в либку
const code = ` 
var id_cnt = 0;
var user_id_arr = [%s];
var users_photos = [];

while (id_cnt < user_id_arr.length) {
    var photos = API.photos.get({"owner_id": user_id_arr[id_cnt], "album_id": "profile"});
    var user_photos = {"user_id": user_id_arr[id_cnt], "photos":[]};
    var photo_cnt = 0;

    while (photo_cnt < photos.count) {
    	var index = (photos.items[photo_cnt].sizes@.type).indexOf("z");
		if (index != -1) {
			user_photos.photos.push(photos.items[photo_cnt].sizes[index].url);
		}	
		photo_cnt = photo_cnt + 1;
    }

    users_photos.push(user_photos);
    id_cnt = id_cnt + 1;
}

return users_photos;
`

func (f *Facade) GetVkUsersPhotosHandler(ctx context.Context, users []uint64) (usersPhotos []model.UserPhotos, err error) {
	vkUsersPhotos, err := f.Vk.GetUsersPhotos(users)
	if err != nil {
		return
	}
	usersPhotos, err = f.downloadUsersPhoto(vkUsersPhotos)
	return
}

func (f *Facade) downloadUsersPhoto(vkUsers []vk.UserPhotos) (usersPhotos []model.UserPhotos, err error) {
	var wg sync.WaitGroup
	wg.Add(len(vkUsers))

	buffer := make(chan model.UserPhotos, len(vkUsers))
	errs := make(chan error, len(vkUsers))

	for _, vkUser := range vkUsers {
		vkUserC := vkUser
		go func() {
			defer wg.Done()
			f.downloadUserPhoto(vkUserC, buffer, errs)
		}()
	}

	wg.Wait()
	usersPhotos = make([]model.UserPhotos, len(vkUsers))
	for i := 0; i < len(vkUsers); i++ {
		err = <-errs
		usersPhotos[i] = <-buffer
		if err != nil {
			return
		}
	}
	return
}

func (f *Facade) downloadUserPhoto(user vk.UserPhotos, userBuffer chan model.UserPhotos, userErrs chan error) {
	var wg sync.WaitGroup
	wg.Add(len(user.Photos))
	var userPhotos model.UserPhotos
	var err error

	userPhotos.UserID = user.UserID
	userPhotos.Photos = make([][]byte, 0, len(user.Photos))

	buffer := make(chan []byte, len(user.Photos))
	errs := make(chan error, len(user.Photos))
	for _, photoUrl := range user.Photos {
		photoUrlC := photoUrl
		go func() {
			defer wg.Done()
			f.downloadPhoto(photoUrlC, buffer, errs)
		}()
	}

	wg.Wait()
	for i := 0; i < len(user.Photos); i++ {
		err = <-errs
		if err != nil {
			userBuffer <- userPhotos
			userErrs <- err
			return
		}
		b := <-buffer
		userPhotos.Photos = append(userPhotos.Photos, b)
	}

	userBuffer <- userPhotos
	userErrs <- err
	return
}

func (f *Facade) downloadPhoto(url string, buffer chan []byte, errs chan error) {
	response, err := http.Get(url)
	if err != nil {
		buffer <- nil
		errs <- err
		return
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		buffer <- nil
		errs <- errors.New("error received non 200 response code")
		return
	}
	b, err := io.ReadAll(response.Body)
	buffer <- b
	errs <- err
	return
}
