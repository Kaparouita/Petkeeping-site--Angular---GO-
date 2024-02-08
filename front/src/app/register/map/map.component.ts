import { Component, Output, EventEmitter } from '@angular/core';

import { GeocodingService } from 'src/app/services/geocoding.service';

declare var OpenLayers: any;

@Component({
  selector: 'app-map',
  templateUrl: './map.component.html',
  styleUrls: ['./map.component.css']
})
export class MapComponent {
  @Output() onMap = new EventEmitter<string>();

  // Method to call when the country is selected/changed
  checkMap() {
    this.onMap.emit(this.address); // TODO: Change this to the lat/lon coordinates
  }

  selectedCountry: string = '';
  city: string = '';
  address: string = '';
  lat: number = 100;
  lon: number = 200;
  map: any;
  mapValidationMessage: string = '';
  showMap: boolean = false;

  constructor(private geocodingService: GeocodingService) {}


  async toggleMap() {
    // Replace with appropriate call to geocoding service
    if (!this.city || !this.address || !this.selectedCountry) {
      this.mapValidationMessage = "Invalid address";
      return;
    }

    try {
      const response = await this.geocodingService.getCoordinates(this.selectedCountry, this.city, this.address).toPromise();
      this.processResponse(response);
    } catch (error) {
      console.error(error);
      this.mapValidationMessage = "Error fetching coordinates";
    }
  }

  processResponse(response: any) {
    if (Object.keys(response).length === 0 || !response[0].display_name.includes("Heraklion")) {
      this.mapValidationMessage = "Invalid address or not available for Heraklion";
      this.showMap = false;
    } else {
      this.mapValidationMessage = "Valid address";
      this.lat = parseFloat(response[0].lat);
      this.lon = parseFloat(response[0].lon);
      this.showMap = true;
      this.createMap();
    }
  }

  createMap() {
    if (this.map) {
      this.map.destroy();
    }

    // Create and configure the OpenLayers map here
    // Use this.lat and this.lon for coordinates
  }

}
