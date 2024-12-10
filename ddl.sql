-- Students Table
CREATE TABLE Students (
    student_id SERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    address TEXT NOT NULL,
    date_of_birth DATE NOT NULL
);

-- Courses Table
CREATE TABLE Courses (
    course_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    department_id INT NOT NULL,
    credits INT NOT NULL
);

-- Professors Table
CREATE TABLE Professors (
    professor_id SERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    address TEXT NOT NULL
);

-- Departments Table
CREATE TABLE Departments (
    department_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT NOT NULL
);

-- Enrollments Table
CREATE TABLE Enrollments (
    enrollment_id SERIAL PRIMARY KEY,
    student_id INT NOT NULL REFERENCES Students(student_id),
    course_id INT NOT NULL REFERENCES Courses(course_id),
    enrollment_date DATE NOT NULL
);

-- Teachings Table
CREATE TABLE Teachings (
    teaching_id SERIAL PRIMARY KEY,
    professor_id INT NOT NULL REFERENCES Professors(professor_id),
    course_id INT NOT NULL REFERENCES Courses(course_id)
);
