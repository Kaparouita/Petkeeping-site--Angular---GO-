import { Component } from '@angular/core';
import { TableModule } from 'primeng/table';
import { UserServiceService } from '../../services/user-service.service';
import { User } from '../../models/user.model';
import { AuthService } from 'src/app/services/auth.service';

interface Column {
  field: string;
  header: string;
}

@Component({
  selector: 'app-admin-home',
  templateUrl: './admin-home.component.html',
  styleUrl: './admin-home.component.css'
})
export class AdminHomeComponent {

  constructor(private userService : UserServiceService , private authService : AuthService) {}
  users: User[] = []; 

  ngOnInit(): void {
    this.userService.getKeepers().subscribe(
      (data) => {
        console.log(data);
        this.users = data;
        this.userService.getOwners().subscribe(
          (data) => {
            console.log(data);
            this.users = this.users.concat(data);
          },
          (error) => {
            console.log(error);
          }
        );
      },
      (error) => {
        console.log(error);
      }
    );

    
  }

  deleteUser(user : User){
    if (user.user_type === "owner"){
      this.userService.deleteOwner(user.id).subscribe(
        (data) => {
          console.log(data);
          this.users.splice(this.users.indexOf(user),1);
        },
        (error) => {
          console.log(error);
        }
      );
    }else if (user.user_type === "keeper"){
      this.userService.DeleteKeeper(user.id).subscribe(
        (data) => {
          console.log(data);
          this.users.splice(this.users.indexOf(user),1);
        },
        (error) => {
          console.log(error);
        }
      );
    }
  }


  onLogout() {
    this.authService.logout();
  }

}
