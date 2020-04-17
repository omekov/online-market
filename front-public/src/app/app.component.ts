import { Component } from '@angular/core';
export interface Category {
  id: number
  name: string
  rus_name: string
  create_at: Date
  update_at: Date
  color: string
  origin_id: number
}
@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  title = 'angular'
  categories: Category[] = [
    {
      id: 1,
      name: "Fruit",
      rus_name: "Фрукты",
      create_at: new Date(),
      update_at: new Date(),
      color: "#fefefe",
      origin_id: 1,
    },
    {
      id: 2,
      name: "Vergies",
      rus_name: "Овощи",
      create_at: new Date(),
      update_at: new Date(),
      color: "#e4e4e4",
      origin_id: 1,
    },
    {
      id: 3,
      name: "Meat",
      rus_name: "Мясо",
      create_at: new Date(),
      update_at: new Date(),
      color: "#e3e3e3",
      origin_id: 2,
    },
  ]

  createCategory(category: Category) {
    this.categories.unshift(category)
  }

  deleteCategory(id: number) {
    this.categories = this.categories.filter(c => c.id !== id)
  }
}
