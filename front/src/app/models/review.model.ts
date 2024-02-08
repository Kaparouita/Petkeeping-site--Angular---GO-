export class Review{
    id: number = 0;
    keeper_id: number = 0;
    owner_id: number = 0;
    booking_id: number;
    rating: number;
    comment: string = '';

    constructor(

        booking_id: number = 0,
        rating: number = 0,
        comment: string = ''
    ) {
        this.booking_id = booking_id;
        this.rating = rating;
        this.comment = comment;
    }
    
}