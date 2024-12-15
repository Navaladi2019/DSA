- used for fixed size array, and following two types of operation
    1) range Query
    2) Update a value



if we use preprocessing then range query is o(1) but when we update then the value becomes o(n)
so when we have an update and we do presum then the avarage time for both read and wrire is o(n) which is not good and no need for preprocessing itself


Binary Indexed Tree (Fenwick tree)
    - used for fixed input array and multiple queries of following types
        prefix operations (sum,product,XOR,OR, etc..)
    - It is actually an array but the concept is tree based
    - Requires o(nlogn) processing time and auxillary space
    - Also known as fenwick tree