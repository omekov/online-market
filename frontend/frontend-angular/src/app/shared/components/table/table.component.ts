import { Component, OnInit, Input } from '@angular/core';
interface IColumn {
  name: string;
  data: any[];
}


@Component({
  selector: 'app-table',
  templateUrl: './table.component.html',
  styleUrls: ['./table.component.scss']
})
export class TableComponent implements OnInit {
  @Input() tableData: IColumn[] = [];
  @Input() isNumeric: boolean = false;
  @Input() isLoading: boolean = false;

  constructor() { }

  ngOnInit(): void {
  }

}
