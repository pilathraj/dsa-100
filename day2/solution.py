def findNumber(a, c):
  o = []
  n = len(a)
  for i in range(n-(c-1)):
    if(a[i]==a[i+1]==a[i+2]):
      o.append(a[i])
  return o
a =[4, 5, 5, 5, 3, 8,22,22,22,5, 22,22,23]
print(findNumber(a, 3))
