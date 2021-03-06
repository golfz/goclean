# เขียน Golang แบบ Clean Architecture
โปรเจคนี้เกิดจากการอ่านหนังสือ [Clean Architecture](https://www.amazon.com/Clean-Architecture-Craftsmans-Software-Structure/dp/0134494164) ของลุงบ็อบ(Robert C. Martin) แล้วยังงงๆ เลยมาลองหาทาง implement ให้ออกมาเป็นโค้ดจริงๆ (เขียนด้วย go)

ผมพยายาม implement โดยให้เป็นไปตาม 4 รูปนี้ครับ (จากเนื้อหาในหนังสือ)

![clean architecture 1](https://raw.githubusercontent.com/golfz/goclean/master/clean-01.jpg)


![clean architecture 1](https://raw.githubusercontent.com/golfz/goclean/master/clean-02.png)


![clean architecture 1](https://raw.githubusercontent.com/golfz/goclean/master/clean-03.png)


![clean architecture 1](https://raw.githubusercontent.com/golfz/goclean/master/clean-03.png)


> โปรเจคนี้เป็นการ implement ตามความเข้าใจของผมนะครับ ซึ่งอาจจะไม่ถูกต้องก็ได้ (ซึ่งผมเองก็ยังไม่ค่อยมั่นใจเหมือนกันครับ) หากใครนำไปใช้ก็สามารถนำไปปรับได้ตามความเหมาะสมครับ


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
โดยที่ Service จะรอรับ Request ที่พอร์ต 8081
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
