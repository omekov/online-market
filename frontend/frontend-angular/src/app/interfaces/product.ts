export interface IProduct {
  id: number,
  name: string,
  description: string,
  price: number,
  updated_at: Date,
  created_at: Date,
  stock: IStock,
  category: ICategory,
}

export interface IStock {
  id: number,
  name: string,
  description: string,
  precent: number,
  created_at: Date,
}

export interface ICategory {
  id: number,
  name: string,
  description: string,
  created_at: Date,
}

export interface ISection {
  backgroundImage: string,
  reverse: boolean,
  content: ISectionContent,
}
interface ISectionContent {
  name: string,
  description: string,
  image: string,
  extra: string,
  route: string,
  buttonName: string,
}
