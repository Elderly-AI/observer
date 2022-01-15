package crawler

import (
	"context"
	"errors"
	"fmt"
	"github.com/Elderly-AI/observer/crawler/internal/pkg/model"
	"github.com/go-vk-api/vk"
	"strconv"
	"strings"
)

const MaxVkUsersCountRequest = 75
const MaxVkUserRequestBatchSize = 25
const MaxVkUsersRequestBatchCount = 3

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
		user_photos.photos.push(photos.items[photo_cnt].sizes[index].url);
    	photo_cnt = photo_cnt + 1;
    }

    users_photos.push(user_photos);
    id_cnt = id_cnt + 1;
}

return users_photos;
`

func (f *Facade) GetVkUsersPhotosHandler(ctx context.Context, users []uint64) (urls [][]string, err error) {
	chunks, err := getVkUsersRequestChunks(users)
	if err != nil {
		return
	}
	//TODO gorutine run
}

func getVkUsersRequestChunks(users []uint64) (chunks [][]uint64, err error) {
	if MaxVkUsersCountRequest < len(users) {
		err = errors.New("error too many users in request") // TODO var
		return
	}
	chunks = make([][]uint64, 0)
	for offset := 0; offset*MaxVkUserRequestBatchSize < len(users); offset += MaxVkUserRequestBatchSize {
		chunks = append(chunks, users[offset:min(offset+MaxVkUserRequestBatchSize, len(users))])
	}
	return
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func (f *Facade) getUsersProfilePhotos(usersID []uint64) (usersProfilePhotos []model.VkUser, err error) {
	if len(usersID) > MaxVkUserRequestBatchSize {
		return nil, errors.New("batch size is too big")
	}

	execCode := fmt.Sprintf(code, Uint64ArrToString(usersID))
	err = f.vkClient.CallMethod("execute", vk.RequestParams{
		"code": execCode,
	}, &usersProfilePhotos)
	return usersProfilePhotos, err
}

func Uint64ArrToString(arr []uint64) string {
	b := make([]string, len(arr))
	for i, v := range arr {
		b[i] = strconv.FormatUint(v, 10)
	}
	return strings.Join(b, ",")
}
