import { ComponentFixture, TestBed } from '@angular/core/testing';

import { VisitorFindKeepersComponent } from './visitor-find-keepers.component';

describe('VisitorFindKeepersComponent', () => {
  let component: VisitorFindKeepersComponent;
  let fixture: ComponentFixture<VisitorFindKeepersComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [VisitorFindKeepersComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(VisitorFindKeepersComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
