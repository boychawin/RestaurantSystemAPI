# RestaurantSystemAPI

Restaurant System API

ออกแบบ API ระบบร้านอาหาร

Functions หลัก

- ตั้งค่า
  - User สำหรับพนักงานใช้งานระบบ 💚
  - Product รายการอาหารในเมนู มี ชื่อ รายละเอียด รูปภาพ ราคา เป็นต้น 💚
  - Product Category หมวดหมู่อาหาร เช่น กับข้าว สเต๊ก สลัด อาหารจานเดียว ขนม ของหวาน น้ำ เป็นต้น 💚
  - Table โต๊ะ สำหรับรับรองลูกค้าที่นั่งทานที่ร้าน 💚

- ระบบสั่งอาหาร
  - เปิดโต๊ะ / เปิดบิลได้ 💚
  - สั่งอาหารเข้าในบิลที่เปิดได้ 💚
  - รายการสั่งอาหาร (แต่ละรายการ) มีสถานะ = รอยืนยัน - กำลังเตรียมอาหาร - เสริฟอาหาร - ยกเลิก 💚
  - การสั่งอาหารสั่งได้หลายรอบ โดยรายการอาหารจะแสดงแยกรอบของการสั่ง 💚
  - หากรายการสั่งอาหารยังไม่เสริฟ จะสามารถยกเลิกรายการนั้นได้ 💚

- ระบบรับชำระ / เชคบิล
  - รวมยอดชำระของทั้งบิลได้  💚
  - เก็บข้อมูลจำนวนเงินที่จ่ายและแสดงเงินทอน  💚
  - เมื่อชำระแล้ว เป็นการปิดโต๊ะ / ปิดบิล  💚

- ระบบจองโต๊ะ
  - โต๊ะมีสถานะ = ว่าง - จอง - เปิด - ไม่ใช้งาน 💚
  - สามารถจองโต๊ะได้ หากสถานะเป็น ว่าง 💚
  - การจองโต๊ะต้องใส่ชื่อลูกค้าที่จองได้ 💚

- ระบบสมัครสมาชิก
  - สามารถสมัครสมาชิกได้ โดยเก็บเลขบัตรประชาชน 💚
  - ตอนคิดเงินหากแสดงบัตรประชาชน แล้วเชคได้ว่าเป็นสมาชิกจะได้ส่วนลดท้ายบิล 10% 💚
  - สมาชิกมีวันหมดอายุ หากเลยวันแล้วใช้ส่วนลดไม่ได้ 💚
  - สามารถต่ออายุบัตรได้ 💚

- รายละเอียดอื่น ๆ
  - แต่ละบิล ต้องการเก็บข้อมูลจำนวนลุกค้าในโต๊ะนั้น ๆ 💚
  - แต่ละบิล ต้องการเก็บข้อมูลเพศและช่วงอายุของลูกค้าได้ 💚

- ระบบรายงาน
  - ต้องการแสดงข้อมูลสรุป รายวัน - รานเดือน ของข้อมูลดังนี้ (ยังมีปัญญา) 💔
    - จำนวนรายรับทั้งหมด (ย้อนหลัง 7 วัน, 15 วัน, 1 เดือน, 3 เดือน) 💔
    - จำนวนรายรับ แยกตาม Product Category (ย้อนหลัง 7 วัน, 15 วัน, 1 เดือน, 3 เดือน) 💔
    - จำนวนบิลทั้งหมด แยกตาม Product Category (ย้อนหลัง 7 วัน, 15 วัน, 1 เดือน, 3 เดือน) 💔
    - จำนวนบิลทั้งหมด (ย้อนหลัง 7 วัน, 15 วัน, 1 เดือน, 3 เดือน) 💚
    - จำนวนลูกค้าทั้งหมดที่เข้าร้าน (ย้อนหลัง 7 วัน, 15 วัน, 1 เดือน, 3 เดือน) 💔
    - จำนวนลูกค้าทั้งหมดที่เข้าร้าน แยกตาม ช่วงอายุ (ย้อนหลัง 7 วัน, 15 วัน, 1 เดือน, 3 เดือน) 💚
    - จำนวนลูกค้าทั้งหมดที่เข้าร้าน แยกตาม เพศ (ย้อนหลัง 7 วัน, 15 วัน, 1 เดือน, 3 เดือน) 💔
    - รายการอาหาร top 10 ของร้านในแต่ละเดือน 💚
    - รายชื่อลูกค้าที่เข้าซ้ำและจำนวนครั้งที่เข้าร้านทั้งหมดในช่วง 15 วัน, 1 เดือน, 3 เดือน, 6 เดือน, 12 เดือน 💔

Assignments

- ออกแบบ ER
- ออกแบบ Database schemas
- ออกแบบ RESTFul API
- Implement RESTFul API ของระบบด้วยภาษาที่ถนัด โดยเทสแต่ละ API ผ่าน Postman

<https://dbdiagram.io/d/Restauran-System-Diagram-650176e702bd1c4a5e7ac33a>

## Installation
