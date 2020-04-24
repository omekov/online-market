import { Component, OnInit } from "@angular/core";

@Component({
  selector: "app-menu",
  templateUrl: "./menu.component.html",
  styleUrls: ["./menu.component.scss"],
})
export class MenuComponent implements OnInit {
  isMenuBtn: boolean = false;
  constructor() {}

  ngOnInit(): void {}
  OpenMenu(): void {
    this.isMenuBtn = !this.isMenuBtn;
  }
}
