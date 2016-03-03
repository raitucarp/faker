package faker

import "reflect"

type Phone struct {
	Number string
	Format string
}

/*
 self.phoneNumber = function (format) {
      format = format || faker.phone.phoneFormats();
      return faker.helpers.replaceSymbolWithNumber(format);
  };

  // FIXME: this is strange passing in an array index.
  self.phoneNumberFormat = function (phoneFormatsArrayIndex) {
      phoneFormatsArrayIndex = phoneFormatsArrayIndex || 0;
      return faker.helpers.replaceSymbolWithNumber(faker.definitions.phone_number.formats[phoneFormatsArrayIndex]);
  };

  self.phoneFormats = function () {
    return faker.random.arrayElement(faker.definitions.phone_number.formats);
  };

*/

func (phone *Phone) Number_(params ...interface{}) string {
	format := phone.Format_()

	types := kindOf(params...)
	if len(types) >= 1 && types[0] == reflect.String {
		format = params[0].(string)
	}

	phone.Number = replaceSymbolWithNumber(format, '#')

	return phone.Number
}

func (phone *Phone) Format_() string {
	format := getData("Phone", "PhoneFormat")
	phone.Format = sample(format)
	return phone.Format
}
