# 常见格式的tag选项

| 格式 | 用途 | 用于 |
| --- | --- | --- |
| json | used by the encoding/json package | detailed at json.Marshal() |
| xml | used by the encoding/xml package | detailed at xml.Marshal() |
| bson | used by gobson | detailed at bson.Marshal() |
| protobuf | used by github.com/golang/protobuf/proto | detailed int the package doc |
| yaml | used by the gopkg.in/yaml.v2 package | detailed at yaml.Marshal() |
| db | used by the github.com/jmoiron/sqlx package also used by github.com/go-gorp/gorp package | |
| orm | used by the github.com/astaxie/beego/orm package | detailed at Models - Beego ORM |
| gorm | used by the github.com/jinzhu/gorm package | examples can be found in their doc: Models |
| valid | used by the github.com/asaskevich/govalidator package | examples can be found in the project page |
| datastore | used by appengine/datastore (Google App Engine platform, Datastore service) | detailed at Properties |
| schema | used by github.com/gorilla/schema to fill a struct with HTML form values | detailed in the package doc |
| asn | used by the encoding/asn1 package | detailed at asn1.Marshal() and asn1.Unmarshal() |