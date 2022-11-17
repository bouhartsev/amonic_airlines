const REQUIRED_FIELD = 'Обязательно для заполнения'

export const firstNameValidation = {
    required: REQUIRED_FIELD,
    validate: (value: string) => {
        if (value.match(/[а-яА-Я]/)) {
            return 'Имя не может содержать русские буквы'
        }
        return true;
    }
}

export const lastNameValidation = {
    required: REQUIRED_FIELD,
    validate: (value: string) => {
        if (value.match(/[а-яА-Я]/)) {
            return 'Фамилия не может содержать русские буквы'
        }
        return true;
    }
}

export const passportNumberValidation = {
    required: REQUIRED_FIELD,
    // validate: (value: string) => {
    //     if (value.length < 10) {
    //         return 'Passport number should be equel 10 numbers'
    //     }
    //     return true;
    // }
}

export const phoneValidation = {
    required: REQUIRED_FIELD,
    // validate: (value: string) => {
    //     if (value.length < 10) {
    //         return 'Passport number should be equel 10 numbers'
    //     }
    //     return true;
    // }
}