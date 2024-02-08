import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { UserServiceService} from '../../services/user-service.service';
import { AuthService } from 'src/app/services/auth.service';
import { LoginResp } from 'src/app/models/login.model';

@Component({
  selector: 'app-admin-login',
  templateUrl: './admin-login.component.html',
  styleUrl: './admin-login.component.css'
})
export class AdminLoginComponent {
  username: string = '';
  password: string = '';

  constructor(private userService : UserServiceService,private router : Router,private authService : AuthService) {}

  onLogin() {
    let params = new LoginResp(this.username, this.password, '', 0);

    this.authService.loginadmin(params).subscribe(
      (data) => {
        console.log(data);
        let url = '/admin-home/';
        console.log(url);
        this.router.navigate([url]);
      },
      (error) => {
        console.log(error);
      }
    );
  }
}
