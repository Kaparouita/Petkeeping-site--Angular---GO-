export class Booking{
    id: number = 0;
    owner_id: number;
    keeper_id: number;
    pet_id: number;
    start_date: string;
    end_date: string;
    price: number;
    status: string;
    message:  string = '';

    constructor(
        owner_id: number = 0,
        keeper_id: number = 0,
        pet_id: number = 0,
        start_date: string = '',
        end_date: string = '',
        price: number = 0,
        status: string = ''
    ) {
        this.owner_id = owner_id;
        this.keeper_id = keeper_id;
        this.pet_id = pet_id;
        this.start_date = start_date;
        this.end_date = end_date;
        this.price = price;
        this.status = status;
    }
    
}