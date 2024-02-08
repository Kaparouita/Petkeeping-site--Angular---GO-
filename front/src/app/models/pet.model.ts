export class Pet{
    id: number = 0;
    pet_id: string;
    owner_id: number;
    name: string;
    age: number;
    type: string;
    breed: string;
    gender: string;
    birth_year: number;
    weight: number;
    description: string;
    photo: string;

    constructor(
        pet_id: string = '',
        owner_id: number = 0,
        name: string = '',
        age: number = 0,
        type: string = '',
        breed: string = '',
        gender: string = '',
        birth_year: number = 0,
        weight: number = 0,
        description: string = '',
        photo: string = ''
    ){
        this.pet_id = pet_id;
        this.owner_id = owner_id;
        this.name = name;
        this.age = age;
        this.type = type;
        this.breed = breed;
        this.gender = gender;
        this.birth_year = birth_year;
        this.weight = weight;
        this.description = description;
        this.photo = photo;
    }
}