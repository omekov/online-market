import { NgModule } from "@angular/core";
import { CommonModule } from "@angular/common";
import { CartComponent } from "./components/cart/cart.component";
import { HeaderComponent } from "./layouts/header/header.component";
import { MainComponent } from "./layouts/main/main.component";
import { BannerComponent } from "./layouts/header/banner/banner.component";
import { MenuComponent } from "./layouts/header/menu/menu.component";
import { RouterModule } from "@angular/router";
import {
  NotLayoutComponent,
  LayoutComponent,
} from "./layouts/layout.component";

@NgModule({
  declarations: [
    CartComponent,
    HeaderComponent,
    MainComponent,
    BannerComponent,
    MenuComponent,
    NotLayoutComponent,
    LayoutComponent,
  ],
  imports: [CommonModule, RouterModule],
  exports: [
    HeaderComponent,
    MenuComponent,
    LayoutComponent,
    NotLayoutComponent,
  ],
})
export class SharedModule {}
