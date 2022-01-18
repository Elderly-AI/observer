package vk

const MaxVkUsersCountRequest = 75
const MaxVkUserRequestBatchSize = 25

const code = ` 
var id_cnt = 0;
var user_id_arr = [%s];
var users_photos = [];

while (id_cnt < user_id_arr.length) {
    var photos = API.photos.get({"owner_id": user_id_arr[id_cnt], "album_id": "profile"});
    var user_photos = {"user_id": user_id_arr[id_cnt], "photos":photos.items};
 	users_photos.push(user_photos);
    id_cnt = id_cnt + 1;
}

return users_photos;
`
