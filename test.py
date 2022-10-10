import psycopg2

print('connecting in progress')

try:
    connection = psycopg2.connect(
        database = 'std_1690_apl',
        user = 'std_1690_apl',
        password = 'porolotbdapl',
        port = '322',
        host = 'std-mysql.ist.mospolytech.ru'
    )
    print("[INFO] connection established")
    cur = connection.cursor()
except:
    print("[INFO] no connection with database")