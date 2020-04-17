import { CategoryFormComponent } from './category-form.component';
import { FormBuilder } from '@angular/forms';

describe('CategoryFormComponent', () => {
  let component: CategoryFormComponent;

  beforeEach(() => {
    component = new CategoryFormComponent(new FormBuilder);
  });

  it('Форме 4 контроллера', () => {
    expect(component.categoryForm.contains('name')).toBeTruthy()
    expect(component.categoryForm.contains('rus_name')).toBeTruthy()
    expect(component.categoryForm.contains('color')).toBeTruthy()
    expect(component.categoryForm.contains('origin_id')).toBeTruthy()
  });
  it('Форме 4 контроллерах валидация', () => {
    const ctrlName = component.categoryForm.get('name')
    ctrlName.setValue('')
    expect(ctrlName.valid).toBeFalsy()
    const ctrlRusname = component.categoryForm.get('rus_name')
    ctrlRusname.setValue('')
    expect(ctrlRusname.valid).toBeFalsy()
    const ctrlColor = component.categoryForm.get('color')
    ctrlColor.setValue('')
    expect(ctrlColor.valid).toBeFalsy()
    const ctrlOriginId = component.categoryForm.get('origin_id')
    ctrlOriginId.setValue('')
    expect(ctrlOriginId.valid).toBeFalsy()

  })
});
