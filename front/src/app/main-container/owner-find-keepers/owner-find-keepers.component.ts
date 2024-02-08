import { Component } from '@angular/core';
import { UserServiceService } from 'src/app/services/user-service.service';
import { ActivatedRoute } from '@angular/router';
import { User } from 'src/app/models/user.model';
import { Pet } from 'src/app/models/pet.model';
import { Booking } from 'src/app/models/booking.model';
import { DatePipe } from '@angular/common';

@Component({
  selector: 'app-owner-find-keepers',
  templateUrl: './owner-find-keepers.component.html',
  styleUrl: './owner-find-keepers.component.css'

})
export class OwnerFindKeepersComponent {

  ownerPets: Pet[] = [];
  petOptions: string[] = [];
  keepers: User[] = [];
  orderByOptions = [
    { label: 'CatPrice', value: 'cat_price' },
    { label: 'DogPrice', value: 'dog_price' },
    { label: 'Duration', value: 'duration' },
    { label: 'Distance', value: 'distance' }
  ];

  hasRequested: boolean = false;
  userId: number = 0;
  visible: boolean = false;
  message: string = '';

  selectedOrder: string = 'Select an order';
  selectedPet: Pet = new Pet();
  selectedPetName: string = '';

  startDate: Date = new Date();
  endDate: Date = new Date();
  diffDays: number = 1;

  constructor(private userService : UserServiceService,private route: ActivatedRoute,private datepipe : DatePipe ) {}

  ngOnInit(): void {
    this.route.paramMap.subscribe(params => {
      const userId = params.get('user_id') ; // convert id to a number
      this.userId = parseInt(userId || '0', 10);

      this.userService.getAvailableKeepers(parseInt(userId || '0', 10)).subscribe(
        data => {
          this.keepers = data; 
        },
        error => {
          console.error('Error fetching owner data', error);
        }
      );

      this.userService.getOwner(parseInt(userId || '0', 10)).subscribe(
        data => {
          this.ownerPets = data.pets; 
          for (let pet of this.ownerPets) {
            this.petOptions.push(pet.name);
          }
        },
        error => {
          console.error('Error fetching owner data', error);
        }
      );
      this.userService.getOwnerBookings(parseInt(userId || '0', 10)).subscribe(
        data => {
          for (let booking of data) {
            if (booking.status === 'requested') {
              this.hasRequested = true;
              break;
            }
          }
        },
        error => {
          console.error('Error fetching owner data', error);
        }
      );
    });
  }

  onOrderChange() {
    
    switch (this.selectedOrder) {
      case 'cat_price':
        this.orderKeepersByPrice('cat');
        break;
      case 'dog_price':
        this.orderKeepersByPrice('dog');
        break;
      case 'duration':
        this.userService.getOrderedKeepers(this.userId,this.selectedOrder).subscribe(
          data => {
            this.keepers = data; 
          },
          error => {
            console.error('Error fetching owner data', error);
          }
        );
        break;
      case 'distance':
        this.userService.getOrderedKeepers(this.userId,this.selectedOrder).subscribe(
          data => {
            this.keepers = data; 
          },
          error => {
            console.error('Error fetching owner data', error);
          }
        );
        break;
      default:
        break;
    }
  }

  orderKeepersByPrice(type : string) {
    this.keepers.sort((a, b) => {
      if (type === 'cat') {
        if (a.cat_keep === false){
          a.cat_price = 999;
        }
        if (b.cat_keep === false){
          b.cat_price = 999;
        }
        return a.cat_price - b.cat_price;
      }
      else {
        if (a.dog_keep === false){
          a.dog_price = 999;
        }
        if (b.dog_keep === false){
          b.dog_price = 999;
        }
        return a.dog_price - b.dog_price;
      }
    });
  }

  onDateChange() {
    if (this.startDate && this.endDate) {
        const diffTime = this.endDate.getTime() - this.startDate.getTime();
        if (diffTime < 0) {
           alert('End date must be after start date');
        }
        const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24)); 
        this.diffDays = diffDays;
    }
  }

  book(keeperId: number,dogPrice: number,catPrice: number) {
    if (this.hasRequested) {
      alert('ERROR : You have already requested a booking');
      return;
    }
    if (this.selectedPetName === '') {
      alert('Please select a pet');
      return;
    }
    var price = 0;
    if (this.selectedPet.type === 'dog') {
      price = dogPrice * this.diffDays;
    }
    else {
      price = catPrice * this.diffDays;
    }

    const book = new Booking(this.userId,keeperId,this.selectedPet.id,this.formatDate(this.startDate),this.formatDate(this.endDate),
    price);
    book.message = this.message;
    this.userService.createBooking(book).subscribe(
      data => {
        alert('Booking successful');
        this.visible = false;
      },
      error => {
        console.error('Error booking keeper', error);
        alert('Error booking keeper');
      }
    );
  }

  onPetChange() {
    for (let pet of this.ownerPets) {
      if (pet.name === this.selectedPetName) {
        this.selectedPet = pet;
      }
    }
    console.log(this.selectedPet);
  }

  showDialog() {
      this.visible = true;
  }

  formatDate(date: Date): string {
    let formatBirthday = this.datepipe.transform(
      date,
      'yyyy-MM-ddTHH:mm:ssZZZZZ'
    );

    return formatBirthday as string;
  }

}
