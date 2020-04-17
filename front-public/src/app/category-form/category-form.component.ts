import { Component, OnInit, Output, EventEmitter } from '@angular/core';
import { Category } from '../app.component';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

@Component({
  selector: 'app-category-form',
  templateUrl: './category-form.component.html',
  styleUrls: ['./category-form.component.scss']
})
export class CategoryFormComponent implements OnInit {
  @Output() onCreate = new EventEmitter<Category>()
  public categoryForm: FormGroup
  constructor(private fb: FormBuilder) {
    this.categoryForm = fb.group({
      name: ['', Validators.required],
      rus_name: ['', Validators.required],
      color: ['', Validators.required],
      origin_id: [1, Validators.required],
    })
   }
  ngOnInit(): void {
  }
  Create(): void {
    this.onCreate.emit(this.categoryForm.value)
    this.categoryForm.reset()
  }
}
