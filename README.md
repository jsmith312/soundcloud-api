## Examples
In main.go:
```
client := soundcloud.NewClient('YOUR_CLIENT_ID',
                               'YOUR_CLIENT_SECRET',
                               'USERNAME',
                               'PASSWORD')

```
### GetUser()
With the client authenticated, we can get [User](https://github.com/jsmith312/go-soundcloud-api/tree/master/soundcloud/user) data:
```
user := client.GetUser()
```
### GetUserGroups()
Get the [Groups](https://github.com/jsmith312/go-soundcloud-api/tree/master/soundcloud/group) that the user has joined
```
groups := client.GetUserGroups()
```
### GetTracks()
Get the [Tracks](https://github.com/jsmith312/go-soundcloud-api/tree/master/soundcloud/track) the user has uploaded
```
tracks := client.GetTracks()
```
### RemoveFromGroup(GroupID, TrackID)
Remove the associated track from the group
```
client.GetUserGroups(GroupID, TrackID)
```
