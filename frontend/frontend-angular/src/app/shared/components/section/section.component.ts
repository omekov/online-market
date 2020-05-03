import { Component, OnInit, Input } from '@angular/core';
import { ISection } from 'src/app/interfaces/product';

@Component({
  selector: 'app-section',
  templateUrl: './section.component.html',
  styleUrls: ['./section.component.scss']
})
export class SectionComponent implements OnInit {
  @Input() section: ISection;

  constructor() { }

  ngOnInit(): void {
  }

}
