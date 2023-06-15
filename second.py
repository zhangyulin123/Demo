import sys
import addressbook_pb2

person = addressbook_pb2.Person()
person.name = "Water"
person.id = 1234
person.email = "water@example.com"

# 向 Person 对象中添加两个 PhoneNumber 对象
phone1 = person.phones.add()
phone1.number = "17612132841"
phone1.type = addressbook_pb2.Person.HOME
phone2 = person.phones.add()
phone2.number = "13521495637"
phone2.type = addressbook_pb2.Person.WORK

# 创建一个新的 AddressBook 对象
address_book = addressbook_pb2.AddressBook()
address_book.people.append(person)

# 将 AddressBook 对象序列化为二进制文件
with open("addressbook.bin", "wb") as f:
    f.write(address_book.SerializeToString())

# 读取二进制文件并反序列化为 AddressBook 对象
with open("addressbook.bin", "rb") as f:
    address_book_data = f.read()
    new_address_book = addressbook_pb2.AddressBook()
    new_address_book.ParseFromString(address_book_data)

# 打印新的 AddressBook 对象
print("新的AddressBooK对象为："new_address_book)
