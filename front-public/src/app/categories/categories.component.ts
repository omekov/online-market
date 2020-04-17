import { Component, OnInit, Input, Output, EventEmitter } from '@angular/core';
import { Category } from '../app.component';

@Component({
  selector: 'app-categories',
  templateUrl: './categories.component.html',
  styleUrls: ['./categories.component.scss']
})
export class CategoriesComponent implements OnInit {
  @Input() category: Category
  @Output() onDelete = new EventEmitter<number>()

  constructor() { }

  ngOnInit(): void {
  }
  Delete(id) {
    this.onDelete.emit(id)
  }

}
