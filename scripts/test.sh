set -e 

echo "--------- SINGLE TESTS ------------" 

echo "get_result json" | nc -u localhost 2000 -w 5
echo "get_result xml" | nc -u localhost 2000 -w 5
echo "get_result native" | nc -u localhost 2000 -w 5
echo "get_result proto" | nc -u localhost 2000 -w 5
echo "get_result avro" | nc -u localhost 2000 -w 5
echo "get_result yaml" | nc -u localhost 2000 -w 5
echo "get_result msg_pack" | nc -u localhost 2000 -w 5

echo "--------- SINGLE TESTS DONE --------------" 
echo 
echo "--------- MULTICAST TEST ----------------" 

echo "get_result all" | nc -u localhost 2000 -w 10

echo "--------- MULTICAST TEST DONE -------------" 