import time

start = time.time()
for i in range(1000000):
    print("hello cik yg ke -", i)

end = time.time()
print("waktu : ", end - start, "detik")