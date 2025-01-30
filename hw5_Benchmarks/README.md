# Финальные результаты запуска тестов
| Бенчмарк                                   |  |  |
|--------------------------------------------|-----------------------|--|
| BenchmarkAddWithChannel-8                  | 3868                  | 287756 ns/op |
| BenchmarkAddWithMutexChunked-8             | 4104                  | 286935 ns/op |
| BenchmarkAddWithWaitGroup-8                | 1736                  | 671965 ns/op |
| BenchmarkSerializeJSON1KB-8                | 681540                | 1708 ns/op |
| BenchmarkSerializeJSON1MB-8                | 711                   | 1490305 ns/op |
| BenchmarkSerializeJSON10MB-8               | 69                    | 16675923 ns/op |
| BenchmarkSerializeGob1KB-8                 | 537688                | 2119 ns/op |
| BenchmarkSerializeGob1MB-8                 | 4605                  | 264124 ns/op |
| BenchmarkSerializeGob10MB-8                | 836                   | 1491821 ns/op |
| BenchmarkSerializeMsgPack1KB-8             | 2131112               | 513.6 ns/op |
| BenchmarkSerializeMsgPack1MB-8             | 10826                 | 97000 ns/op |
| BenchmarkSerializeMsgPack10MB-8            | 1549                  | 736997 ns/op |
| BenchmarkStandardSort100-8                 | 963997                | 1208 ns/op |
| BenchmarkBubbleSort100-8                   | 163350                | 6781 ns/op |
| BenchmarkStandardSort1000-8                | 73796                 | 16395 ns/op |
| BenchmarkBubbleSort1000-8                  | 1455                  | 804774 ns/op |
| BenchmarkCountSubstringsStringsCount-8     | 260139                | 4450 ns/op |
| BenchmarkCountSubstringsRegex-8            | 234224                | 5020 ns/op |
| BenchmarkCountSubstringsManual-8           | 265263                | 4671 ns/op |