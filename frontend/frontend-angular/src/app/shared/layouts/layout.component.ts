import { Component } from "@angular/core";

@Component({
  selector: "app-layout",
  template: `
    <app-header></app-header>
    <app-main>
      <router-outlet></router-outlet>
    </app-main>
  `,
})
export class LayoutComponent {}

@Component({
  selector: "app-not-layout",
  template: `
    <router-outlet></router-outlet>
  `,
})
export class NotLayoutComponent {}

