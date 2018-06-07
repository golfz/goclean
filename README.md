# เขียน Golang แบบ Clean Architecture

## TODO
- [x] GET
- [ ] POST
- [ ] PUT
- [ ] DELETE


## Example
โค้ดนี้สามารถทดสอบได้ง่ายๆเลย ด้วยการรันด้วยคำสั่ง 
```bash
go run main.go
```
โดยที่ Service ของเราจะรอรับ Request ที่พอร์ต 8081
จะทดสอบส่ง request ด้วย Browser หรือจะใช้ [Postman](https://www.getpostman.com) ก็ได้เช่นกัน


### GET
#### URI : 
```
/example/{key}
```
#### ตัวอย่าง Request
```
http://localhost:8081/example/3
```

#### ตัวอย่าง Response
```json
{
    "key": "3",
    "name": "Example product",
    "color": "Red",
    "remain": 213,
    "created-date": "2018-06-07",
    "url": "/example/2",
    "rel": "self"
}
```
