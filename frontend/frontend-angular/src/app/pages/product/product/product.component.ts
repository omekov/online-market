import { Component, OnInit } from '@angular/core';
import { ISection } from 'src/app/interfaces/product';

@Component({
  selector: 'app-product',
  templateUrl: './product.component.html',
  styleUrls: ['./product.component.scss']
})
export class ProductComponent implements OnInit {

  constructor() { }
  sections: ISection[] = [
    {
      backgroundImage: "assets/img/sections/section-bg-1.svg",
      reverse: true,
      content: {
        name: "Beautiful food & takeaway, delivered to your door.",
        description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500.",
        image: "assets/img/sections/section-1.png",
        route: "/order",
        buttonName: "Place an Order",
        extra: "4.8 out of 5 based on 2000+ reviews"
      },
    },
    {
      backgroundImage: "assets/img/sections/section-bg-2.svg",
      reverse: true,
      content: {
        name: "The home of fresh products",
        description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500.",
        image: "assets/img/sections/section-2.png",
        route: "/company",
        buttonName: "Learn about us",
        extra: null
      },
    },
    {
      backgroundImage: "",
      reverse: false,
      content: {
        name: "Order online with our simple checkout.",
        description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500.",
        image: "assets/img/sections/section-3.png",
        route: "/faq",
        buttonName: "See our FAQ",
        extra: null
      },
    },
  ];
  ngOnInit(): void {
    console.log('test')
  }

}
