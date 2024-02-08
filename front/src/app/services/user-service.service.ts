import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

import { User } from '../models/user.model';
import { LoginResp } from '../models/login.model';
import { Pet } from '../models/pet.model';
import { Booking } from '../models/booking.model';
import { Message } from '../models/message.model';
import { Review } from '../models/review.model';

@Injectable({
  providedIn: 'root',
})
export class UserServiceService {
  constructor(private http: HttpClient) {}

  saveOwner(params: User) {
    return this.http.post<User>(`http://127.0.0.1:3000/users/owners/`, params);
  }

  deleteOwner(id: number) {
    return this.http.delete(`http://127.0.0.1:3000/users/owners/` + id);
  }

  addPet(params: Pet){
    return this.http.post<Pet>(`http://127.0.0.1:3000/pets/`, params);
  }

  saveKeeper(params: User) {
    return this.http.post<User>(`http://127.0.0.1:3000/users/keepers/`, params);
  }

  getOwner(id: number) {
    return this.http.get<User>(`http://127.0.0.1:3000/users/owners/` + id);
  }

  getOwners() {
    return this.http.get<User[]>(`http://127.0.0.1:3000/users/owners/`);
  }

  getAvailableKeepers(id : number) {
    return this.http.get<User[]>(`http://127.0.0.1:3000/users/owners/` + id + `/findKeepers`);
  }

  getOrderedKeepers(id : number, orderBy : string) {
    return this.http.get<User[]>(`http://127.0.0.1:3000/users/owners/` + id + `/orderKeepers?orderBy=` + orderBy);
  }

  getKeeper(id: number) {
    return this.http.get<User>(`http://127.0.0.1:3000/users/keepers/` + id);
  }

  updateOwner(params: User) {
    return this.http.put<User>(`http://127.0.0.1:3000/users/owners/`+ params.id, params);
  }

  updateKeeper(params: User) {
    return this.http.put<User>(`http:////127.0.0.1:3000/users/keepers/`+ params.id, params);
  }

  getAdmin(id: number) {
    return this.http.get<User>(`http:////127.0.0.1:3000/users/admins/` + id);
  }

  getKeepers() {
    return this.http.get<User[]>(`http:////127.0.0.1:3000/users/keepers/`);
  }

  DeleteKeeper(id: number) {
    return this.http.delete(`http:////127.0.0.1:3000/users/keepers/` + id);
  }

  GetNumberOfCats() {
    return this.http.get<number>(`http:////127.0.0.1:3000/admin/stats`);
  }

  GetNumberOfDogs() {
    return this.http.get<number>(`http:////127.0.0.1:3000/admin/stats`);
  }

  GetBookingsByKeeperId(id: number) {
    return this.http.get<Booking[]>(`http:////127.0.0.1:3000/users/keepers/` + id + `/bookings`);
  }

  UpdateBooking(params: Booking) {
    return this.http.put<Booking>(`http:////127.0.0.1:3000/bookings/` +params.keeper_id ,params);
  }

  createBooking(book : Booking) {
    return this.http.post<Booking>(`http:////127.0.0.1:3000/bookings/`, book);
  }

  getOwnerBookings(id: number) {
    return this.http.get<Booking[]>(`http:////127.0.0.1:3000/users/owners/` + id + `/bookings`);
  }

  GetReviewsByKeeper(id: number) {
    return this.http.get<User>(`http:////127.0.0.1:3000/users/keepers/` + id + `/reviews`);
  }

  sendMessage(message : Message) {
    return this.http.post<Message>(`http:////127.0.0.1:3000/messages`, message);
  }

  getMessagesByName(name : string) {
    return this.http.get<Message[]>(`http:////127.0.0.1:3000/messages?name=` + name);
  }
  GetBooking( id: number) {
    return this.http.get<Booking>(`http:////127.0.0.1:3000/bookings/` + id);
  }

  createMessage(message : Message) {
    return this.http.post<Message>(`http:////127.0.0.1:3000/messages`, message);
  }

  CreateReview(review : Review) {
    return this.http.post<Review>(`http:////127.0.0.1:3000/reviews/` ,review);
  }

  GetPetKeepersDays(id: number) {
    return this.http.get<number>(`http:////127.0.0.1:3000/users/keepers/` + id + `/petKeepersDays`);
  }

  GetPetKeepersBookings(id: number) {
    return this.http.get<number>(`http:////127.0.0.1:3000/users/keepers/` + id + `/bookingsNumber`);
  }

}
