import { Component } from '@angular/core';
import { User } from 'src/app/models/user.model';
import { Router } from '@angular/router';
import { UserServiceService } from 'src/app/services/user-service.service';
import { ActivatedRoute } from '@angular/router';
import { AuthService } from 'src/app/services/auth.service';
import { DatePipe } from '@angular/common';

interface Days{
  days: number;
}

interface Bookings{
  bookings: number;
}

@Component({
  selector: 'app-keepers-review',
  templateUrl: './keepers-review.component.html',
  styleUrl: './keepers-review.component.css'
})
export class KeepersReviewComponent {

  user = new User();
  owners : User[] = [];
  reviews : any;
  days: number = 0;
  bookings: number = 0;
  constructor( private authService : AuthService,private userService : UserServiceService,private route: ActivatedRoute, private datePipe: DatePipe) { }

  ngOnInit(): void {
    this.route.paramMap.subscribe(params => {
      this.userService.getOwners().subscribe(
        data => {
          this.owners = data; // Assign the emitted value to user
        },
        error => {
          console.error('Error fetching owner data', error);
        }
      );


      const userId = params.get('user_id') ; // convert id to a number

      this.userService.GetReviewsByKeeper(parseInt(userId || '0', 10)).subscribe(
        data => {
          this.reviews = data; // Assign the emitted value to user
        },
        error => {
          console.error('Error fetching keeper data', error);
        }
      );

      this.userService.GetPetKeepersDays(parseInt(userId || '0', 10)).subscribe(
        data => {
          this.days = data; // Assign the emitted value to user
        },
        error => {
          console.error('Error fetching keeper data', error);
        }
      );

      this.userService.GetPetKeepersBookings(parseInt(userId || '0', 10)).subscribe(
        data => {
          this.bookings = data; // Assign the emitted value to user
        },
        error => {
          console.error('Error fetching keeper data', error);
        }
      );
    });
    
  }


  getOwnerName(id: number) {
    return this.owners.find(owner => owner.id === id)?.first_name;
  }
}
