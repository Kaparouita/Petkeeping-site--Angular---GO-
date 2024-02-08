import { ComponentFixture, TestBed } from '@angular/core/testing';

import { OwnerAddPetComponent } from './owner-add-pet.component';

describe('OwnerAddPetComponent', () => {
  let component: OwnerAddPetComponent;
  let fixture: ComponentFixture<OwnerAddPetComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [OwnerAddPetComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(OwnerAddPetComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
