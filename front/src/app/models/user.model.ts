import { Pet } from "./pet.model";

export class User {
    id: number = 0 ;
    username: string;
    email: string;
    password: string;
    first_name: string;
    last_name: string;
    birth_date: string;
    gender: string;
    user_type: string;
    country: string;
    city: string;
    address: string;
    url: string;
    job: string;
    phone: string;
    space_type: string = '';
    cat_keep: boolean = false;
    dog_keep: boolean = false;
    cat_price: number = 8;
    dog_price: number = 8;
    space_desc: string = '';
    lat: number = 0;
    lon: number = 0;
    pets: Pet[] = [];

    constructor(
        username: string = '',
        email: string = '',
        password: string = '',
        first_name: string = '',
        last_name: string = '',
        birth_date: string = '',
        gender: string = '',
        user_type: string = '',
        country: string = '',
        city: string = '',
        address: string = '',
        url: string = '',
        job: string = '',
        phone: string = '',
        lat: number = 0,
        lon: number = 0,
    ) {
        this.username = username;
        this.email = email;
        this.password = password;
        this.first_name = first_name;
        this.last_name = last_name;
        this.birth_date = birth_date;
        this.gender = gender;
        this.user_type = user_type;
        this.country = country;
        this.city = city;
        this.address = address;
        this.url = url;
        this.job = job;
        this.phone = phone;
        this.lat = lat;
        this.lon = lon;
    }
}