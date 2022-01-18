package vk

const MaxVkUsersCountRequest = 75
const MaxVkUserRequestBatchSize = 25

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
