import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

interface petCount{
  cats: number;
  dogs: number;
}

interface userCount{
  owners: number;
  keepers: number;
}

interface money{
  keeper: number;
  app: number;
}

@Injectable({
  providedIn: 'root'
})
export class AdminService {

  constructor(private http: HttpClient) {}

  getPetCount() {
    return this.http.get<petCount>(`http://127.0.0.1:3000/users/admin/petsNumber`);
  }

  getUserCount() {
    return this.http.get<userCount>(`http://127.0.0.1:3000/users/admin/usersNumber`);
  }

  getMoney() {
    return this.http.get<money>(`http://127.0.0.1:3000/users/admin/money`);
  }
}
