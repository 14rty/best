import mysql.connector as myc

print('connecting in progress')

connect_args ={
    'host': 'std-mysql.ist.mospolytech.ru',
    'port': '3306',
    'user': 'std_1690_apl',
    'password': 'porolotbdapl',
}

try:
    db = myc.connect(**connect_args)
    cursor = db.cursor(named_tuple=True)
    print("[INFO] connection established")

except:
    print("[INFO] no connection with database")