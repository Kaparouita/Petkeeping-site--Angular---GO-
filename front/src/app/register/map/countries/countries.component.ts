import { Component, Output, EventEmitter } from '@angular/core';
@Component({
  selector: 'app-countries',
  templateUrl: './countries.component.html',
  styleUrl: './countries.component.css'
})
export class CountriesComponent {
  @Output() countryChange = new EventEmitter<string>();
  selectedCountry: string = 'Greece';

  // Method to call when the country is selected/changed
  selectCountry() {
    this.countryChange.emit(this.selectedCountry);
  }
}
