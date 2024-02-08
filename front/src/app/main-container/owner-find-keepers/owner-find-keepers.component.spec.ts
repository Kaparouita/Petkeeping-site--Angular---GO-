import { ComponentFixture, TestBed } from '@angular/core/testing';

import { OwnerFindKeepersComponent } from './owner-find-keepers.component';

describe('OwnerFindKeepersComponent', () => {
  let component: OwnerFindKeepersComponent;
  let fixture: ComponentFixture<OwnerFindKeepersComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [OwnerFindKeepersComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(OwnerFindKeepersComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
