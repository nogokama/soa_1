set -e 

echo "get_result json" | nc -u localhost 2000 -w 2
echo "get_result xml" | nc -u localhost 2000 -w 2
echo "get_result native" | nc -u localhost 2000 -w 2
echo "get_result proto" | nc -u localhost 2000 -w 2
echo "get_result avro" | nc -u localhost 2000 -w 2
echo "get_result yaml" | nc -u localhost 2000 -w 2
echo "get_result msg_pack" | nc -u localhost 2000 -w 2