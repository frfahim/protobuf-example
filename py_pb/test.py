import employee_pb2 as employee_pb

employee = employee_pb.Employee()
employee.id = 1
employee.name = "Fahim"
employee.designation = "Software Engineer"
employee.salary = 100000.0
employee.SerializeToString()
sd = employee.SerializeToString()
print(sd)
employee2 = employee_pb.Employee()
employee2.ParseFromString(sd)
print(employee2)