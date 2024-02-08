import { Component } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute } from '@angular/router';
import { UserServiceService } from 'src/app/services/user-service.service';

import { Pet } from '../../models/pet.model';

@Component({
  selector: 'app-owner-add-pet',
  templateUrl: './owner-add-pet.component.html',
  styleUrl: './owner-add-pet.component.css'
})
export class OwnerAddPetComponent {
  addPetForm: FormGroup | any;
  petTypes: any[] | undefined; 

  owner_id: number | undefined;
  constructor(private fb: FormBuilder,private route: ActivatedRoute,private userService: UserServiceService) { }

  ngOnInit() {
    // Generate a random 10-digit string
    const randomString = Math.floor(Math.random() * 9000000000) + 1000000000;
    const randomStringWithDigits = randomString.toString();
    console.log(randomStringWithDigits);
      
    // Initialize the form
    this.route.paramMap.subscribe(params => {
      const userId = params.get('user_id') ;
      this.owner_id = parseInt(userId || '0', 10);
    });

    this.addPetForm = this.fb.group({
      pet_id : [randomStringWithDigits],
      owner_id: [this.owner_id, Validators.required],
      name: ['', Validators.required],
      age: [0, Validators.required],
      type: ['', Validators.required], // Dropdown
      breed: ['', Validators.required],
      gender: ['', Validators.required],
      birth_year: [0, Validators.required],
      weight: [0, Validators.required],
      description: [''],
      photo: [''] // Optional
    });

  }


  onSubmit() {
    if (this.addPetForm.valid) {
      const formValues = this.addPetForm.value;
      var pet = new Pet(
        formValues.pet_id,
        formValues.owner_id,
        formValues.name,
        parseInt(formValues.age || '0', 10),
        formValues.type,
        formValues.breed,
        formValues.gender,
        parseInt(formValues.birth_year || '0', 10),
        parseInt(formValues.weight || '0', 10),
        formValues.description,
        formValues.photo
        )
      this.userService.addPet(pet).subscribe(
        (data) => {
          console.log(data);
          alert('Pet added successfully');
        },
        (error) => {
          console.log(error);
          alert('Error adding pet' + error.message);
        }
      );
    }
    else {
      console.log('Invalid form');
    }
  }
}

