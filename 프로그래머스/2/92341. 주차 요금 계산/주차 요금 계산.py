def solution(fees, records):
    fee_minute = fees[0]
    fee = fees[1]
    minute = fees[2]
    minute_fee = fees[3]
    
    cars = dict()
    car_fee = dict()
    
    for record in records:
        input = record.split()
        time = input[0]
        car_num = input[1]
        
        if car_num in cars:
            if len(cars[car_num]) == 0:
                cars[car_num].append(time)
            else:
                prev_h, prev_m = cars[car_num][0].split(':')
                prev_total = int(int(prev_h) * 60) + int(prev_m)
                now_h, now_m = time.split(":")
                now_total = int(int(now_h) * 60) + int(now_m)
                time = now_total - prev_total
                cars[car_num].pop()
                
                
                if car_num in car_fee:
                    car_fee[car_num].append(time)
                else:
                    car_fee[car_num] = []
                    car_fee[car_num].append(time)
        else:
            cars[car_num] = []
            cars[car_num].append(time)

            

    
    for car in cars:
        if len(cars[car]) == 0:
            continue
        prev_h, prev_m = cars[car][0].split(':')
        prev_total = int(int(prev_h) * 60) + int(prev_m)
        now_total = 23 * 60 + 59
        total_time = int(now_total - prev_total)
        cars[car].pop()
        if car in car_fee:
            car_fee[car].append(total_time)
        else:
            car_fee[car] = []
            car_fee[car].append(total_time)

    answer = []
    
    print(car_fee)
    
    for car in car_fee:
        total_fee = 0
        total_minute = 0
        for m in car_fee[car]:
            total_minute += m
        
        
        if total_minute <= fee_minute:
            total_fee += fee
        else:
            total_minute = total_minute - fee_minute
            total_fee += fee
            추가_분 = total_minute / minute
            if total_minute % minute != 0 :
                추가_분 = int(추가_분) + 1
            total_fee += 추가_분 * minute_fee
        
        answer.append((car, total_fee))

    
    answer.sort()
    
    real_answer = []
    
    for i in answer:
        real_answer.append(int(i[1]))
        
    return real_answer