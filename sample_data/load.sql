load data local infile '~/go/src/github.com/sunyatsuntobee/server/sample_data/organizations.csv' into table organizations 
character set utf8 
fields terminated by ','  optionally 
enclosed by '"' escaped by '"' 
ignore 1 lines;


load data local infile '~/go/src/github.com/sunyatsuntobee/server/sample_data/activities.csv' into table activities 
character set utf8 
fields terminated by ',' 
enclosed by '"' escaped by '"' 
ignore 1 lines;


load data local infile '~/go/src/github.com/sunyatsuntobee/server/sample_data/stages.csv' into table activity_stages 
character set utf8 
fields terminated by ','  optionally 
enclosed by '"' escaped by '"' 
ignore 1 lines;
