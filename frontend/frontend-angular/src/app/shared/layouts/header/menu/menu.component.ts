import { Component, OnInit, HostListener } from "@angular/core";

@Component({
  selector: "app-menu",
  templateUrl: "./menu.component.html",
  styleUrls: ["./menu.component.scss"],
})
export class MenuComponent implements OnInit {
  isMenuBtn: boolean = false;
  constructor() {}

  @HostListener("window:resize", ["$event"])
  onResize(event) {
    if (event.target.innerWidth === 991) {
      this.isMenuBtn = true
    }
  }
  ngOnInit(): void {}
  OpenMenu(): void {
    this.isMenuBtn = !this.isMenuBtn;
  }
}
