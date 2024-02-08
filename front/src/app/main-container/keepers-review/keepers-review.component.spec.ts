import { ComponentFixture, TestBed } from '@angular/core/testing';

import { KeepersReviewComponent } from './keepers-review.component';

describe('KeepersReviewComponent', () => {
  let component: KeepersReviewComponent;
  let fixture: ComponentFixture<KeepersReviewComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [KeepersReviewComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(KeepersReviewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
