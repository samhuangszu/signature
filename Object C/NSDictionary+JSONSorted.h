#import <Foundation/Foundation.h>
@interface NSDictionary (JSONSorted) 
- (NSArray *)sortedKeys;
- (NSString *)sortedJSONString;
@end