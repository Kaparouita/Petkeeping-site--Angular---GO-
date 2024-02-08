import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { BehaviorSubject, Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { LoginResp } from '../models/login.model';
import { Router } from '@angular/router';

@Injectable({
  providedIn: 'root'
})

export class AuthService {
  curr_user = 0;

  constructor(private http: HttpClient, private router: Router) {
  }

  public get currentUserValue() {
    return this.curr_user;
  }

  login(params: LoginResp) {
    return this.http.post<LoginResp>(`http://127.0.0.1:3000/login/`, params)
      .pipe(map(user => {
        // Assuming the response contains a token
        if (user && (user.user_id != 0)) {
          this.curr_user = user.user_id;
        }
        return user;
      }));
  }

  loginadmin(params: LoginResp) {
    return this.http.post<LoginResp>(`http://127.0.0.1:3000/login/admin/`, params)
      .pipe(map(user => {
        // Assuming the response contains a token
        if (user && (user.user_id != 0)) {
          this.curr_user = user.user_id;
        }
        return user;
      }));
  }

  logout() {
    this.curr_user = 0;
    this.router.navigate(['/login']);
  }

  isAuthenticated() {
    return this.curr_user != 0;
  }
}

