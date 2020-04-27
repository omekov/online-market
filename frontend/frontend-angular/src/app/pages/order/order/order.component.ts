import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-order',
  templateUrl: './order.component.html',
  styleUrls: ['./order.component.scss']
})
export class OrderComponent implements OnInit {

  constructor() { }

  ngOnInit(): void {
  }
  categories = [
    {
      id: 1,
      name: "Full Menu",
      active: true,
    },
    {
      id: 2,
      name: "Burgers",
      active: false,
    },
    {
      id: 3,
      name: "Sides",
      active: false,
    },
    {
      id: 4,
      name: "Drinks",
      active: false,
    }
  ]
  filterOrder(id: number): void {
    return
  }
  addCart() {

  }
  products = [
    {
      id: 1,
      name: "Drink Fig & Lime",
      description: "Lorem ipsum dolor sit amet consectetur adipisicing elit.",
      price: 4,
      url: "/assets/img/lime.png"
    },
    {
      id: 2,
      name: "Drink Fig & Lime",
      description: "Blanditiis doloribus fugiat pariatur facere accusantium.",
      price: 4,
      url: "/assets/img/lime.png"
    },
    {
      id: 2,
      name: "Drink Fig & Lime",
      description: "Blanditiis doloribus fugiat pariatur facere accusantium.",
      price: 4,
      url: "/assets/img/lime.png"
    },
    {
      id: 2,
      name: "Drink Fig & Lime",
      description: "Blanditiis doloribus fugiat pariatur facere accusantium.",
      price: 4,
      url: "/assets/img/lime.png"
    },
    {
      id: 2,
      name: "Drink Fig & Lime",
      description: "Blanditiis doloribus fugiat pariatur facere accusantium.",
      price: 4,
      url: "/assets/img/lime.png"
    },
    {
      id: 2,
      name: "Drink Fig & Lime",
      description: "Blanditiis doloribus fugiat pariatur facere accusantium.",
      price: 4,
      url: "/assets/img/lime.png"
    }
  ]
}
