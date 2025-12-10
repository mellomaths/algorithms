import unittest
import time

from random import randint


def binary_search(buckets: list[float], item: float) -> float:
    bot = 0
    top = len(buckets) - 1
    while bot <= top:
        mid = (top + bot) // 2
        guess = buckets[mid]
        if guess == item:
            return mid
        if guess > item:
            top = mid - 1
        elif guess < item:
            bot = mid + 1
    return None


class BinarySearchTests(unittest.TestCase):
    
    def test_multiple_random_lists(self):
        
        def get_item_from_list(buckets: list, item: float) -> float:
            try:
                return buckets.index(item)
            except ValueError:
                return None
        
        def random_test_case():
            item = randint(0, 1000)
            buckets = list(set([randint(0, 1000) for _ in range(1000)]))
            buckets.sort()
            self.assertEqual(binary_search(buckets, item), get_item_from_list(buckets, item))
        
        """
        Test 100 random lists and items.
        """
        for _ in range(100):
            random_test_case()
    
    def test_fixed_lists(self):
        self.assertEqual(binary_search([1,3,5,7,9], 3), 1)
        self.assertIsNone(binary_search([1,3,5,7,9], 8))
        
    def test_execution_time(self):
        start_time = time.time()
        n = 1000000
        item = binary_search([x for x in range(n)], randint(1, n))
        print(f"\n--- {time.time() - start_time} seconds for {n} elements ---")
        self.assertIsNotNone(item)

if __name__ == "__main__":    
    unittest.main(verbosity=2) 
    
    
    