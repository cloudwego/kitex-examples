namespace go order

service OrderService {
    void Create(1: string userId, 2: string commodityCode, 3: i32 count)
}
