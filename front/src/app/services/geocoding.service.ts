import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class GeocodingService {

  constructor(private http: HttpClient) {}

  getCoordinates(country: string, city: string, address: string) {
    const url = 'https://forward-reverse-geocoding.p.rapidapi.com/v1/forward';
    const params = {
      street: address,
      city: city,
      country: country,
      "accept-language": "en",
      polygon_threshold: "0.0"
    };
    const headers = {
      "X-RapidAPI-Key": "54711c88damsh1f9e11569b8ab7ep1f42f3jsn646e18dae2d6",
      "X-RapidAPI-Host": "forward-reverse-geocoding.p.rapidapi.com"
    };
    return this.http.get(url, { params, headers });
  }
}
