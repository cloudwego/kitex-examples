namespace go storage

service StorageService {
    void Deduct(1: string commodityCode, 2: i32 count)
    i32 Calculate(1: string commodityCode, 2: i32 count)
}