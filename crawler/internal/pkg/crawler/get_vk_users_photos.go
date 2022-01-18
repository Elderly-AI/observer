package crawler

import (
	"context"
	"errors"
	"github.com/Elderly-AI/observer/crawler/internal/pkg/model"
	"io"
	"net/http"
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

func (f *Facade) downloadUsersPhoto(vkUsers []model.VkUser) (usersPhotos []model.UserPhotos, err error) {
	usersPhotos = make([]model.UserPhotos, 0, len(vkUsers))
	for _, vkUser := range vkUsers {
		var userPhotos model.UserPhotos
		userPhotos, err = f.downloadUserPhoto(vkUser)
		if err != nil {
			return
		}
		usersPhotos = append(usersPhotos, userPhotos)
	}
	return
}

func (f *Facade) downloadUserPhoto(vkUser model.VkUser) (userPhotos model.UserPhotos, err error) {
	userPhotos.UserID = vkUser.UserID
	userPhotos.Photos = make([][]byte, 0, len(vkUser.Photos))
	for _, photoUrl := range vkUser.Photos {
		var b []byte
		b, err = f.downloadPhoto(photoUrl)
		if err != nil {
			return
		}
		userPhotos.Photos = append(userPhotos.Photos, b)
	}
	return
}

func (f *Facade) downloadPhoto(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return nil, errors.New("error received non 200 response code")
	}
	b, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}
