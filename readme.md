
# groupbycolumn

Group by column is a comamnd line utility that takes csv's and prints out unique rows in the following format.
```
Input:

Bob   2017-01-02
Bob   2017-03-02
Bob   2017-03-02
Bob   2017-02-01
Cindy 2017-02-01
Cindy 2017-02-03

Output:

Bob
	2017-01-02
	2017-02-01
	2017-03-02
Cindy
	2017-02-01
	2017-02-03
```
